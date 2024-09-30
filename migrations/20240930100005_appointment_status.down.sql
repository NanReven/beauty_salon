ALTER TABLE appointments
    ALTER COLUMN status TYPE VARCHAR(255);

DROP TYPE appointment_status;