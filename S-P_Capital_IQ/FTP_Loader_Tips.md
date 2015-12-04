# S&P Capital IQ

S&P Capital IQ provide access to their data in two main forms:

* API 
* FTP + Loader Application + Database

If you are using large volumes of data then they will steer you away from the API to the latter option. This guide applies to the 'FTP' and doesn't reference the API at all.

## The FTP Option

The first thing you should know about the FTP option is that is more complicated than the API.

The setup requires you to host a copy of the Capital IQ database. They will specify the hardware, operating system and database software you need to use (MS SQL or Oracle are the only options). The setup requires two machines, one to host the database server and another to run the 'loader' software which downloads the daily deltas and updates the database. There is a little flexibility in what hardware and software they will approve for use, but not much. If you are looking to keep startup costs down, then the additional servers and licenses should be kept in mind.

You might be thinking at this point that you would simply use the Capital IQ database as your primary repository of data but if performance is important to your application, then you should dismiss that idea now. The Capital IQ data is fully normalised and contains millions of records, it is _slow_, with many common queries requiring upwards of half a dozen joins. Pulling out only the data you actually require and storing it in some other fashion whether that's plain text or another database may be more suitable for your use case. 

### Data Sets

The most important advice I could give at this point is to establish your exact data requirements down to the last detail with S&P in advance. All the Capital IQ data is split into major data sets within which there are multiple add-on packages which provide additional data and metadata. You cannot assume that the basic packages contain the necessary information because they probably will not. To avoid going back again and again to request further packages, all of which will require renegotiation of your contract, you should know exactly what you will and won't be given access to, down to the exact data point, before heading to the integration stage.

#### Coverage

Some data just isn't available from the Database.

* The **Global Instruments Cross Reference Service (GICRS)** which maps Capital IQ internal IDs to standard company or stock identifiers such as the **ISIN**, **CINS**, **CUSIP** or **Valor** is a standalone file which you must parse and store yourself.

* Transcripts are only available via an API.

### Documentation

S&P Capital IQ provide a lot of documentation, too much and yet not enough. The documentation takes the form of many PDFs which cannot be searched from the support portal. There's plenty giving high level descriptions of packages which isn't nearly descriptive or specific enough about the data items provided and then there is schema documentation showing the relationships between the various tables. What is missing is much in the way of useful examples or a quick start guide. Finding the specific information you require will often prove difficult and ultimately, if you're like us, you'll just end up diving into the database to try to figure it all out for yourself. Be prepared to set aside plenty of time for this and ask questions of your support representative, that's what they are there for.



