Public API:
    // health check
    http://localhost:9090/api/student/v1/public/health

    // Get all students (GET method)
    http://localhost:9090/api/student/v1/public/students

    // Get one student by id (GET method)
    http://localhost:9090/api/student/v1/public/student?id=value

    // Get student(s) with key words (PATCH method)
    http://localhost:9090/api/student/v1/public/studentWithKeywords


Private API (need token):
    // Add a student (POST method)
    http://localhost:9090/api/student/v1/staff/student

    // Update Email of a student (PUT method) by id
    http://localhost:9090/api/student/v1/staff/student

    // Delete a student (DELETE method) by id
    http://localhost:9090/api/student/v1/staff/student