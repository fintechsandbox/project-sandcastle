# Morningstar::Datawarehouse

Morningstar data warehouse crawler implemented in Ruby. Data is loaded from the Morningstar FTP and transfered into the Amazon S3.

The gem also provides the mechanisms for load back transfered data from S3, parse it and send to specify appenders (see usage section below). 

## Installation

Add this line to your application's Gemfile:

```ruby
gem 'morningstar-datawarehouse'
```

And then execute:

    $ bundle

## Configuration

Place somewhere the following section for crawler configuration:

```ruby
Morningstar::Datawarehouse::Configuration.configure do |config|
  # Credentials for accessing Morningstar FTP
  config.ftp_source = 'ftp.morningstar.com'
  config.ftp_username = 'user'
  config.ftp_password = 'pwd'

  # AWS S3 credentials for transfering Morningstar data to.  
  config.aws_access_key = 'AWS_ACCESS_KEY_ID'
  config.aws_access_secret = 'AWS_ACCESS_SECRET_KEY'
  config.aws_region  = 'AWS_REGION'
  config.aws_s3_bucket  = 'AWS_S3_BUCKET'
  
  
  config.logger = ::Logger.new('morningstar_data_warehouse.log')
end
```


## Usage

**Loading Morningstar daily data**

```ruby
Morningstar::Datawarehouse::DataTransfuser.new("Daily/DataWarehouse", "my_s3_bucket").transfer(Date.today.prev_day.strftime("%Y%m%d"))
```

**Loading Morningstar monthly data**

```ruby
Morningstar::Datawarehouse::DataTransfuser.new("Monthly/DataWarehouse", "my_s3_bucket").transfer(Date.today.strftime("%Y%m"))
```

**Process transferred data**

```ruby
# Implement appenders for process parsed data
class MyAppender < Morningstar::Datawarehouse::Parsers::DataAppender

    def append(data)
    #process data here
    end
  
    #Note: 'shares' is array of SharesClassId 
    def deleted_shares(shares=[], universe=nil)
      #process delete shares
    end
    
    #Note: 'shares' is array of SharesClassId 
    def under_review_shares(shares=[], universe=nil)
      #process shares with 'Under Review' status
    end
    
    #Note: 'shares' is array of SharesClassId 
    def updated_shares(shares=[], universe=nil)
      #process updated shares
    end
end

# Process all daily and monthly data
Morningstar::Datawarehouse::DataLoader.new([MyAppender.new]).process_data()

# Process monthly data
Morningstar::Datawarehouse::DataLoader.new([MyAppender.new]).process_monthly_data(Date.today.strftime("%Y%m"))

# Process daily data
Morningstar::Datawarehouse::DataLoader.new([MyAppender.new]).process_daily_data(Date.today.prev_day.strftime("%Y%m%d"))

```