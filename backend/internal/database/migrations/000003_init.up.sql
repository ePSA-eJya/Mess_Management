CREATE TYPE semester_type AS ENUM ('ODD', 'EVEN');

CREATE TABLE semesters(
                          semester_id INT PRIMARY KEY ,
                          academic_year VARCHAR(100) NOT NULL ,
                          semester_type semester_type  NOT NULL DEFAULT 'ODD',
                          start_date DATE ,
                          end_date DATE
);