# Expense Tracker API  

This project is a simple yet powerful Expense Tracker API built using **Rust**, **Axum**, **SQLx**, and **Tokio**. It leverages **PostgreSQL** as the database to efficiently manage expenses.  

## Features  
- Create, read, update and delete expenses  
- Secure authentication and authorization  
- Optimized performance with async Rust  
- Easy deployment with Docker  

## Installation  

1. Clone the repository:  
   ```sh  
   git clone https://github.com/yourusername/expense-tracker-api.git  
   cd expense-tracker-api  
   ```  
2. Install dependencies:  
   ```sh  
   cargo build  
   ```  
3. Set up the database and run migrations:  
   ```sh  
   DATABASE_URL=postgres://user:password@localhost/expense_tracker_api sqlx migrate run  
   ```  
4. Start the server:  
   ```sh  
   cargo run  
   ```  

## Usage  

Once the server is running, you can access the API at `http://localhost:8000`. Use tools like **Postman** or **cURL** to test endpoints.  

## Contributing  

Contributions are welcome! Feel free to open issues or submit pull requests. Every contributors is appreciated.  

## License  

This project is licensed under the MIT License.  


