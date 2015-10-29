# Kyper Data
FactSet has an interesting dataset called [Revere](http://www.factset.com/campaigns/revere) that contain data on supply chain relationships, geographic revenue breakdowns, and detailed industry/sector information for companies. The data are delivered as a large collection of CSVs, or an install script that builds a Postgres database. They're structured as relational tables, but a graph format/structure may be more useful.

# Elsen POV:
Factset contains similar data to Bloomberg, but makes it more easily available (and cheaper) by delivering it as CSV through FTP. 

##### [factset.sql](factset.sql)
Sample Factset schema derived from factset CSVs delivered via Factset's ftp.
