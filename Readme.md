# Mutual Fund Data ETL Assignment

### I have used Go programming language & PostgreSQL as the database

### The code structure is as follows:
- dailyLoad (This script is responsible for fetching incremental data)
- initialLoad (This script is responsible for fetching historical data, i.e., from 1st April 2006)
- transform (This manager is used to parse the text data, and form data rows out of it)
- load (This manager is used to load the data into the MF_data table)

### Flow 
- For historical data:
    - The data is fetched from the API in batches of 90 days.
    - Once the data is parsed and rows are formed, they are passed to the load manager.
    - Lastly when data is loaded, the timeStamp is updated in timeStampDB.
- For incremental data:
    - The data is fetched from the API, where the start date is fetched from the timeStampDB and end date is currentDate-1.
    - Now same steps are followed, like the historical data.

### Data model
- The table schema is specified in the init.sql file.

### For incremental data
In order to fetch the incremental data, we can form a shell script where-in we can give commands for two cron jobs:
  - The first job will run only once (i.e For the initial data loading)
  - The second job will be set to run every day at midnight 00:00 am (to fetch incremental data of the previous day)

### Completed Tasks Checklist
- The pipeline will fetch all the historical data & incremental data
- I have tried to handle the maximum failures that could happen in the pipeline.
- The initial load got completed in less than 3hrs on my machine.
- To tune the performance, I have used an index.
- I haven't containerized the project.
