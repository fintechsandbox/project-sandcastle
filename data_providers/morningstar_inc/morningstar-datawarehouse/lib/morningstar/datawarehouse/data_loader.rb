require 'stringio'
require 'logger'

class StringIO
  def path
  end
end

module Morningstar
  module Datawarehouse
    class DataLoader
      def initialize(data_appenders)
        @data_appenders = data_appenders
        config = Morningstar::Datawarehouse::Configuration.options
        @logger = config.logger
        @s3_bucket = config.aws_s3_bucket
        @client = Aws::S3::Client.new(
            credentials: Aws::Credentials.new(config.aws_access_key, config.aws_access_secret),
            region: config.aws_region)
      end

      def logger
        @logger ||= ::Logger.new('morningstar_data_warehouse.log')
      end

      # This method do next structure
      # Daily/Monthly return hash
      # Universe (SO, FO etc) return hash
      # Date (20151222, 20151223 etc) return array of files
      def set_group_object(list, group)
        # Not very good solution
        list.contents.each do |object|
          keys = object.key.split("/")
          if(keys[0] == "MorningStar")
            group[keys[1]] ||= {}
            group[keys[1]][keys[3]] ||= {}
            group[keys[1]][keys[3]][object.key[/\d{8}/]] ||= []
            group[keys[1]][keys[3]][object.key[/\d{8}/]].push(object)
          end
        end
      end

      def parse_group(group, doc, parser, delta, universe, date)
        s3_objects = group[delta][universe][date]
        zip = nil
        docs = []

        s3_objects.each do |s3_object|
          if ["zip", "gzip", "gz"].include? s3_object.key.split('.').last
            zip = s3_object
          else
            docs.push s3_object
          end
        end

        s3_extractor = Morningstar::Datawarehouse::Extractors::S3Extractor.new(@s3_bucket, @client)
        action_hash = {}

        docs.each do |doc|
          key = doc.key

          action = /(delete|update|underReview)/.match(key).to_s
          action_hash[action] = []
          string_list = ""

          s3_extractor.translate_object(key) do |chunk|
            string_list+= chunk
          end

          action_hash[action] = string_list.split("\r\n")
        end

        s3_extractor.translate_object(zip.key) do |chunk|
          parser.parse chunk
        end if zip

        doc.with_universe(universe).with_deleted_shares(action_hash["delete"])
            .with_under_review_shares(action_hash["underReview"])
            .with_updated_shares(action_hash["update"]).close()

        logger.info  "#{date} Ended \n"
      end


      def process_data
        list =  @client.list_objects({bucket: @s3_bucket})
        group = {}

        doc = Morningstar::Datawarehouse::Parsers::DataXmlDocument.new.with_data_appenders(@data_appenders)

        parser = Nokogiri::HTML::SAX::Parser.new(doc)
        set_group_object(list, group)

        unless group['Monthly'].nil?
          group['Monthly'].keys.each do |universe|
            group['Monthly'][universe].keys.each do |date|
              logger.info  "Started #{universe} #{date} \n"
              parse_group(group, doc, parser, 'Monthly', universe, date)
            end
          end
        end

        group['Daily'].keys.each do |universe|
          group['Daily'][universe].keys.each do |date|
            logger.info  "Started #{universe} #{date} \n"
            parse_group(group, doc, parser, 'Daily', universe, date)
          end
        end
      end

      # If there need of working with one universe only
      def process_universe_data(universe, appenders = [])
        list =  @client.list_objects({bucket: @s3_bucket})
        group = {}
        doc = Morningstar::Datawarehouse::Parsers::DataXmlDocument.new.with_data_appenders(appenders || @data_appenders)
        parser = Nokogiri::HTML::SAX::Parser.new(doc)
        set_group_object(list, group)

        unless group['Monthly'][universe].nil?
          group['Monthly'][universe].keys.each do |date|
            logger.info  "Started #{universe} #{date} \n"
            parse_group(group, doc, parser, 'Monthly', universe, date)
          end
        end

        unless group['Daily'][universe].nil?
          group['Daily'][universe].keys.each do |date|
            logger.info  "Started #{universe} #{date} \n"
            parse_group(group, doc, parser, 'Daily', universe, date)
          end
        end
      end

      def process_daily_data(date)
        list =  @client.list_objects({bucket: @s3_bucket})
        group = {}
        doc = Morningstar::Datawarehouse::Parsers::DataXmlDocument.new.with_data_appenders(@data_appenders)
        parser = Nokogiri::HTML::SAX::Parser.new(doc)
        set_group_object(list, group)

        group['Daily'].keys.each do |universe|
          logger.info "Started #{universe} #{date} \n"
          parse_group(group, doc, parser, 'Daily', universe, date)
        end
      end

      def process_monthly_data(_date)
        list =  @client.list_objects({bucket: @s3_bucket})
        group = {}
        doc = Morningstar::Datawarehouse::Parsers::DataXmlDocument.new.with_data_appenders(@data_appenders)
        parser = Nokogiri::HTML::SAX::Parser.new(doc)
        set_group_object(list, group)

        group['Monthly'].keys.each do |universe|
          group['Monthly'][universe].keys.each do |date|
            # TODO: Parse only previous month
            logger.info "Started #{universe} #{date}\n"
            parse_group(group, doc, parser, 'Monthly', universe, date) if !date.nil? && date.include?(_date)
          end
        end
      end
    end
  end
end

