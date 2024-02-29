using Pkg
Pkg.add("DataFrames")
Pkg.add("Logging")

# Load the required packages
using DataFrames
using Logging

# Set up logging configuration
Logging.configure(level=Logging.Info)

# Function to generate and log sample employee data
function generate_and_log_employee_data(num_employees=10)
    employees = DataFrame(
        emp_id = 1:num_employees,
        department = rand(["IT", "HR", "Sales", "Operations"], num_employees),
        salary = rand(10000:70000, num_employees),
        working_mode = rand(["Hybrid", "Onsite"], num_employees)
    )

    # Log the generated employee data
    @info(employees)
end
