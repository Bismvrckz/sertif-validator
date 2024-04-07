CREATE TABLE IF NOT EXISTS
    tkbai_data (
        id INT PRIMARY KEY AUTO_INCREMENT,
        test_id VARCHAR(255),
        name VARCHAR(255),
        student_number VARCHAR(255),
        major VARCHAR(255),
        date_of_test TIMESTAMP,
        toefl_score VARCHAR(255),
        insert_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );

