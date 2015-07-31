# Elsen's Experience
Alexandria uses the content and context of news stories to analyze the information in real time and provide a sentiment analysis. Using probabilistic algorithms developed from a concept-based, non-linear statistical methodology of testing finance professionals, Alexandria has the ability to model complex dependencies between various words of semantic information, which helps it better identify sentiment than standard rules-based approaches. Sentiment data over a certain period is aggregated into a log net sentiment metric, which can then be compared with other companies in the market. Alexandria prefers to deliver their data in CSV format over FTP.

Elsen used Alexandria's market sentiment analysis data set to perform a series of client-tailored backtests. Data was delivered through flat files. 

## ingest_alexandria.py
Quick script using psycopg2 to parse and write alexandria files to postgresql tables.
