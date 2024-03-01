CREATE TABLE IF NOT EXISTS
    tkbai_data (
        test_id VARCHAR(255) UNIQUE PRIMARY KEY,
        name VARCHAR(255),
        student_number VARCHAR(255),
        major VARCHAR(255),
        date_of_test TIMESTAMP,
        insert_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );

