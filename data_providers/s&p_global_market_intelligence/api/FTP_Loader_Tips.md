# S&P Capital IQ

S&P Capital IQ provide access to their data in two main forms:

* API 
* FTP

If you are using large volumes of data then they will steer you away from the API to the latter option. This guide applies to the 'FTP' and doesn't reference the API at all.

## The FTP Option

The first thing you should know about the FTP option is that it is more complicated than the API.

There are two options here:

* Take the raw data from the FTP and process it yourself.
* Use the Capital IQ software and database to process and store the data.

With the first option, the data arrives in a flat file format with delimited values. This is essentially a non-standard representation of each table in the database with no column data and some additional metadata to indicate whether the line adds, updates or removes a record. There are two different delimiters used depending on whether the data is mostly strings or integers. For a single data set this might be managable, however we took one look and realised we'd spend weeks mapping and parsing every file, nevermind actually understanding the relationships between each 'table'. So we went with the second option.

The second option, requires you run their software which will automatically download, parse and insert the data into a database. They will specify the hardware, operating system and DBMS you need to use (MS SQL or Oracle are the only options). The setup requires two machines, one to host the database server and another to run the 'loader' software which downloads the daily deltas and updates the database. There is a little flexibility in what hardware and software they will approve for use, but not much. If you are looking to keep startup costs down, then the requirement for additional servers and licenses should be kept in mind.

You might be thinking at this point that you would simply use the Capital IQ database as your primary repository of data but if performance is important to your application, then you should dismiss that idea now. The schema is fully normalised and contains millions of records, it is _slow_, with common queries requiring upwards of half a dozen joins. Pulling out only the data you actually require and storing it in some other fashion whether that's plain text, json or another database may be more suitable for your use case. 

### Data Sets

The most important advice I could give at this point is to establish your exact data requirements down to the last detail with S&P in advance. All the Capital IQ data is split into major data sets within which there are multiple add-on packages which provide additional data and metadata. You cannot assume that the basic packages contain the necessary information because they probably will not. To avoid going back again and again to request further packages, all of which will require renegotiation of your contract, you should know exactly what you will and won't be given access to, to the last item, before heading to the integration stage.

As an example - if you want the financial statements you'll be looking at the Latest Financials package. What you may not realise is that this only contains the line items for each quarter. If you want to map those back to the document they came from, whether it's the Income statement, Balance sheet or Cash flow statement, you'll need the 'Financial Display' Add-on. This pattern repeats with often 'basic' metadata abstracted into add-on packages which ultimately you will need to pay more for.

#### Coverage of the FTP

While the S&P coverage, past and present, of companies, securities, estimates, financial and fundamental data is naturally excellant not everything is available using just the FTP alone.

Some data must be imported by other means:

* The **Global Instruments Cross Reference Service (GICRS)** which maps Capital IQ internal IDs to standard company or stock identifiers such as the **ISIN**, **CINS**, **CUSIP** or **Valor** is a standalone file which you must parse and store yourself.

* Transcripts are only available via an API.

### Documentation

S&P Capital IQ provide a lot of documentation, too much and yet not enough. The documentation takes the form of many PDFs which cannot be searched from the support portal. There's plenty giving high level descriptions of packages which isn't nearly descriptive or specific enough about the data items provided and then there is schema documentation showing the relationships between the various tables. What is missing is much in the way of useful examples or a quick start guide. Finding the specific information you require will often prove difficult and ultimately, if you're like us, you'll just end up diving into the database to try to figure it all out for yourself. Be prepared to set aside plenty of time for this and ask questions of your support representative, that's what they are there for.



