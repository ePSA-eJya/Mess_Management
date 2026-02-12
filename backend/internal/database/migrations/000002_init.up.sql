CREATE TYPE admin_type AS ENUM ('OFFICE_IN_CHARGE', 'MESS_IN_CHARGE');

CREATE TABLE admins (
                        id INT PRIMARY KEY,
                        name VARCHAR(100) NOT NULL,
                        admin_type admin_type NOT NULL DEFAULT 'OFFICE_IN_CHARGE',
                        hostel VARCHAR(100) NOT NULL,
                        mess_no INT NOT NULL,
                        phone VARCHAR(15),
                        email VARCHAR(255)
);


-- ------------------------



CREATE TYPE meal_type AS ENUM ('BREAKFAST', 'LUNCH','DINNER');

CREATE TABLE meal_cancellation_records(
    roll INT ,
    semester_id INT ,
    meal_type meal_type NOT NULL DEFAULT 'BREAKFAST',
    date DATE NOT NULL
);


-- ------------------------

CREATE TABLE monthly_bills(
                              bill_id INT PRIMARY KEY ,
                              roll INT NOT NULL,
                              month VARCHAR(100),
                              semester_id INT ,
                              breakfast_count INT,
                              lunch_count INT,
                              dinner_count INT,
                              total_bill FLOAT
);


-- ------------------------

CREATE TABLE semester_bills(
                               bill_id INT PRIMARY KEY ,
                               roll INT NOT NULL,
                               semester_id INT ,
                               total_bill FLOAT
);