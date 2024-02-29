# Install the required package if not already installed
using Pkg
Pkg.add("DataFrames")

# Load the DataFrames package
using DataFrames

# Function to generate sample employee data
function generate_employee_data(num_employees=10)
    employees = DataFrame(
        emp_id = 1:num_employees,
        department = rand(["IT", "HR", "Sales", "Operations"], num_employees),
        salary = rand(10000:70000, num_employees),
        working_mode = rand(["Hybrid", "Onsite"], num_employees)
    )
    return employees
end
