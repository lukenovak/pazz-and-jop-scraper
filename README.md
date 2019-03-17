# pazz-and-jop-scraper
Scrapes the Pazz and Jop polls from 1971-1999 and prints albums and artists then adds them to a database

## Running the scraper
To run the scraper, you first have to configure the approprite database schema, 
with a table called "albums" that has columns for the album and artist names. In addition,
you need to fill in your psql data (I will hopefully be switching this over to environment variables
at some point)

## Planned additions
Adding year to the database, so that I can do data analysis
