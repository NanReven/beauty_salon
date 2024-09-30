CREATE TYPE appointment_status AS ENUM ('pending', 'accepted', 'completed', 'cancelled');

ALTER TABLE appointments
    ALTER COLUMN status TYPE appointment_status USING status::appointment_status;

ALTER TABLE appointments
    ALTER COLUMN status SET DEFAULT 'pending';
