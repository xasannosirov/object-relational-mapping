-- students table for student
CREATE TABLE IF NOT EXISTS students (
    id SERIAL PRIMARY KEY NOT NULL,
    first_name VARCHAR(32) NOT NULL,
    last_name VARCHAR(32) NOT NULL,
    age INT NOT NULL
);

-- courses table for course
CREATE TABLE IF NOT EXISTS courses (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(64) NOT NULL,
    price INT NOT NULL,
    teacher VARCHAR(64) NOT NULL
);

-- student_course table for join courses and students
CREATE TABLE IF NOT EXISTS student_course (
    course_id INT, 
    student_id INT, 
    FOREIGN KEY(course_id) REFERENCES courses(id),
    FOREIGN KEY(student_id) REFERENCES students(id)
);