Welcome to Chrono-Task, a versatile task scheduler written in Go that empowers you to execute commands at specific times. Manage your scheduled tasks effortlessly using a user-friendly configuration file, harnessing the precision and simplicity of Go.
Features:

  Flexible Scheduling: Define tasks and execution times using a simple and intuitive syntax.
  Execution: Run custom commands or scripts at precise intervals.
  Configuration: Easily configure tasks and schedules through the config.yaml file.
  Efficiency: Minimal resource consumption, ideal for various environments.
  Logging: Comprehensive logs and reporting for task executions.

# Getting Started:

Clone the Repository:

    git clone https://github.com/HaziqAliSohail/Chrono-Task.git

Configure Tasks:

 Edit the config.yaml file to define your tasks and their execution schedules.

Build and Run:

    go build
    ./Chrono-Task

Example Configuration (config.yaml):

    tasks:
    - name: "BackupData"
      schedule: "0 2 * * *"
      command: "bash backup.sh"
    - name: "SendReports"
      schedule: "0 18 * * 1-5"
      command: "python send_reports.py"

# Task Schedule Syntax:

The task schedule follows the cron syntax, allowing you to define precise execution times for your tasks. For more information on the cron syntax, refer to Cron Expression.
Contribution:

Contributions and feature requests are welcome! If you encounter any issues or have suggestions for improvements, feel free to open an issue or submit a pull request.

# License:

This project is licensed under the MIT License.
