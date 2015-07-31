#AlphaPack initial yodlee

We're using Yodlee Interactive's Aggregation API. It comes in both SOAP and REST API formats. We use REST.
+https://developer.yodlee.com/Aggregation_API/Aggregation_Quickstart/Aggregation_REST_Quick_Start_Guide

I'd suggest signing up for a free 3month Yodlee trial, this is their "Sandbox" environment. Do not confuse this with the
product being offered through Fintech Sandbox. This is particularly useful because of it gives you access to the
'TestDrive' GUI. This allows you to access all their API functions and visually examine the data returned before writing code.
+https://developer.yodlee.com/TestDrive

Yodlee is different from market-data datasets in that there isn't necessarily a limit to the data you can retrieve. It is only
limited by the actual accounts that a user has.

For the data we receive via the aggregation API we're parsing out information for each of our clients' accounts
eg: credit cards, checking, brokerage
Full list and a high-level description of the data model: https://developer.yodlee.com/Aggregation_API/Aggregation_Services_Guide/Data_Model

General tips: It's important to remember that Yodlee is screen-scraping custodian sites. They can normalize as much they
can but at the end of the day they do not guarantee that all the data will be normalized

Fastlink is their offering to generate an iframe on your site so you don't have to build anything while you are testing out the
data in sample code:
+https://developer.yodlee.com/Aggregation_API/Aggregation_Services_Guide/FastLink_for_Aggregation
