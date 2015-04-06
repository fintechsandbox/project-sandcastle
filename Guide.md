##Table of Contents

About Elsen Sandcastle

Key Terms

Data Providers

Guide

  * Security Master
  * Data Sources
  * Data Series Design
	
Appendix

* Provider Specific Ingest Scripts
  * Alexandria

##About Elsen Sandcastle

Elsen Sandcastle is a collection of guides and scripts that demonstrate how to ingest and develop a database schema that is based on time series data. The scripts are focused on data providers working with the FinTech Sandbox. Each data point, when available, is entered into its own table, with a timestamp when appropriate. It is important to note that Elsen Sandcastle is not meant to be the fastest or best performing database structure; rather, the Sandcastle intends to build a baseline for how data from the FinTech Sandbox can be entered and how it interacts between providers, with a special emphasis on creating a security master in order to identify the data properly.

Elsen Sandcastle is not a representation of Elsen’s internal database structure; rather, it is a reflection of the thought process and experience from working with the FinTech Sandbox data providers.

##Key Terms

Sandcastle – the overall collection of guides and scripts that demonstrate how to ingest and develop a database schema based on time series data

FinTech Sandbox – a collection of partners that provides access to financial market data for financial technology startups to utilize in product development and testing

Dumptruck – the name of the example database created by following Sandcastle. References to Dumptruck are specific for the database as outlined in the guide if followed exactly. 

Security master – a table with all important, necessary identifiers for financial securities that have data within Dumptruck

##Guide

There are several main, crucial components to Dumptruck in order to successfully utilize the data from the FinTech Sandbox. These components are the security master, the source list, and the standard time series schema. The security master will allow information to be properly identified and classified (necessary in order to pull and use data within Dumptruck). The source list will help identify and verify which data providers are providing information (necessary for maintaining validity and reliability of the database, as well as for quickly sourcing and resolving data issues). The standard time series schema allows for all data to be entered in a uniform way (necessary for data visibility as well as querying into the database quickly). 

##Dumptruck structure

In order to create Dumptruck, a schema first needs to be outlined. This schema contains the three necessary components as outlined above. 

Schema.sql

##Security Master

From the schema above, we see the security master start to take shape

Snippet of schema.sql

This shows the structure for our security master. The first, and arguably most crucial, element is the data id (id). This data id is an internal unique identifier, dependent on how you ultimately aggregate data. Some differentiators will be if a provider has data on a strategic business unit or other subcategorized level, such as geographic region, if the holding entity has a different legal identifier. For the purposes of Dumptruck, all securities will be presumed to have standard, unique identifiers. Even with this assumption, it is important to create an internal identifier. This is because although securities may be uniquely identified, there are multiple unique identifiers, such as CUSIP, SEDOL, and exchange/ticker symbol. Creating a security master will tie data identified by one of these indicators to data identified with another indicator representative of the same security. 

This will get more complicated when bringing in data from multiple countries, exchanges, and types of securities. 

[HOW TO INGEST NEW DATA TO TIE TO THIS?????]

##Data Sources


##Data Series Design

As mentioned above, creating a standard time series structure is important for the overall database structure. 

Snippet from schema.sql

In the real world, much like many other things, data providers provide data in all shapes, sizes, and organizations. It will be important to understand these and how they tie into the database structure created for your organization. Here is an example of how data needs to be spliced accordingly to fit it into the structure we have designed.

[Example schema division from current Read Me]

Examples for specific data providers is provided in the appendix.

##Appendix
