CREATE TYPE student_status AS ENUM ('ACTIVE', 'INACTIVE');

CREATE TABLE students (
                          roll INT PRIMARY KEY,
                          name VARCHAR(100) NOT NULL,
                          room_no INT NOT NULL,
                          hostel VARCHAR(100) NOT NULL,
                          mess_no INT NOT NULL,
                          status student_status NOT NULL DEFAULT 'ACTIVE',
                          phone VARCHAR(15),
                          email VARCHAR(255) UNIQUE NOT NULL
);