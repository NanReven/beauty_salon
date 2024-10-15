package service

import (
	"beauty_salon/internal/adapter/repository"
	"beauty_salon/internal/domain/entity"
	"log"
	"os"
	"time"

	"github.com/wneessen/go-mail"
)

type AppointmentService struct {
	appointmentRepo repository.Appointment
	favourRepo      repository.Favour
	masterRepo      repository.Master
}

func NewAppointmentService(appointmentRepo repository.Appointment, favourRepo repository.Favour, masterRepo repository.Master) *AppointmentService {
	return &AppointmentService{appointmentRepo: appointmentRepo, favourRepo: favourRepo, masterRepo: masterRepo}
}

func (uc *AppointmentService) CreateAppointment(userId int, appointment *entity.AppointmentInput) (int, error) {
	if appointment.MasterId <= 0 || len(appointment.Services) == 0 || appointment.AppointmentStart.IsZero() {
		return 0, entity.ErrInvalidAppointmentInput
	}

	appointmentEnd, totalSum, err := uc.getAppointmentInfo(appointment.AppointmentStart.Time, appointment.Services)
	if !uc.CheckMasterAvailability(appointment.AppointmentStart.Time, appointmentEnd, appointment.MasterId) {
		return 0, err
	}

	id, err := uc.appointmentRepo.CreateAppointment(userId, appointment, appointmentEnd, totalSum)
	if err != nil {
		return 0, err
	}

	email, err := uc.masterRepo.GetMasterEmail(appointment.MasterId)
	if err != nil {
		return 0, err
	}
	sendGmail(email)
	return id, nil
}

func (uc *AppointmentService) GetAllAppointments(userId int) ([]entity.AppointmentResponse, error) {
	appointments, err := uc.appointmentRepo.GetAllAppointments(userId)
	if err != nil {
		return appointments, err
	}

	for i := 0; i < len(appointments); i++ {
		services, err := uc.appointmentRepo.GetFavoursByAppointmentId(appointments[i].Id)
		if err != nil {
			return appointments, err
		}
		appointments[i].Services = services
	}
	return appointments, nil
}

func (uc *AppointmentService) GetAppointmentById(userId, appointmentId int) (entity.AppointmentResponse, error) {
	if appointmentId <= 0 {
		return entity.AppointmentResponse{}, entity.ErrInvalidAppointmentInput
	}

	appointment, err := uc.appointmentRepo.GetAppointmentById(userId, appointmentId)
	if err != nil {
		return appointment, err
	}

	services, err := uc.appointmentRepo.GetFavoursByAppointmentId(appointmentId)
	if err != nil {
		return appointment, err
	}

	appointment.Services = services
	return appointment, nil
}

func (uc *AppointmentService) CancelAppointment(userId, appointmentId int) (string, error) {
	if appointmentId <= 0 {
		return "", entity.ErrInvalidAppointmentInput
	}

	appointment, err := uc.appointmentRepo.GetAppointmentById(userId, appointmentId)
	if err != nil {
		return "", err
	} else if appointment.Status == "cancelled" {
		return "", entity.ErrAppointmentCancelled
	}

	return uc.appointmentRepo.CancelAppointment(userId, appointmentId)
}

func (uc *AppointmentService) CheckMasterAvailability(appointmentStart time.Time, appointmentEnd time.Time, masterId int) bool {
	acceptedAppointments, err := uc.appointmentRepo.GetAcceptedAppointments(appointmentStart, masterId)
	if err != nil {
		return false
	}

	log.Println("Input appointment:", appointmentStart, ":", appointmentEnd)
	for _, appointment := range acceptedAppointments {
		log.Println("Accepted appointment:", appointment.AppointmentStart, ":", appointment.AppointmentEnd)
		if appointmentStart.Before(appointment.AppointmentEnd.Time) && appointmentEnd.After(appointment.AppointmentStart.Time) {
			log.Println("Conflict found with appointment:", appointment.AppointmentStart, "-", appointment.AppointmentEnd)
			return false
		}
	}
	return true
}

func (uc *AppointmentService) getAppointmentInfo(appointmentStart time.Time, services []int) (time.Time, float64, error) {
	appointmentEnd := appointmentStart
	var totalSum float64

	for _, id := range services {
		favour, err := uc.favourRepo.GetFavourById(id)
		if err != nil {
			return time.Now(), 0, err
		}
		duration := time.Duration(favour.Duration.Hour()*3600+favour.Duration.Minute()*60) * time.Second
		appointmentEnd = appointmentEnd.Add(duration)
		totalSum += favour.Price
	}
	return appointmentEnd, totalSum, nil
}

func sendGmail(email string) bool {
	message := mail.NewMsg()
	if err := message.From("nanreven@gmail.com"); err != nil {
		log.Printf("failed to set FROM address: %s", err)
		return false
	}
	if err := message.To(email); err != nil {
		log.Printf("failed to set TO address: %s", err)
		return false
	}

	message.Subject("This is my first test mail with go-mail!")
	message.SetBodyString(mail.TypeTextPlain, "This will be the content of the mail.")

	client, err := mail.NewClient("smtp.gmail.com",
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithTLSPortPolicy(mail.TLSMandatory),
		mail.WithPort(587),
		mail.WithUsername(os.Getenv("SMTP_USER")),
		mail.WithPassword(os.Getenv("SMTP_PASSWORD")),
	)

	if err != nil {
		log.Printf("failed to create new mail delivery client: %s", err)
	}
	if err := client.DialAndSend(message); err != nil {
		log.Printf("failed to deliver mail: %s", err)
	}
	return true
}
