require 'nokogiri'

module Morningstar
  module Datawarehouse
    module Parsers
      class XmlParser < Nokogiri::XML::SAX::Document

        attr_accessor :action_hash
        attr_writer :data_appenders

        def initialize
          @object_node = 'InvestmentVehicle'

          @stack_values = []
          @parsed_node = {}

          @data_appenders = []

          @current_node = nil
          @current_value = nil
          @current_hash = nil
        end

        def with_data_appenders(appenders=[])
          @data_appenders += appenders
          self
        end

        def start_element name, attrs = []
          @stack_values.push name
          @current_node = name

          case name
            when 'LegalStructure'
              @parsed_node['legalstructure_id'] = attrs[0][1]
            when 'GrowthOf10KNAV'
              @array = []
              @parsed_node[name] = @array

            when 'TradingExchangeList'
              if @stack_values.include?('TradingExchange')
                @parsed_node['TradingExchange'][name] = @array
                @current_hash = {}
              end

            when 'SEDOLOfficialListingExchangeList'
              @current_hash = {}

            when 'TradingExchange'
              @array = []
              @parsed_node[name] = {}

            when 'HistoryDetail'
              @current_hash = {}
          end
        end

        def characters(string)
          case @current_node
            when 'EndDate'
              @current_hash[@current_node] = string

            when 'Exchange'
              if @stack_values.include?('TradingExchangeList')
                @current_hash[@current_node] = string
              end

            when 'TradingSymbol'
              if @stack_values.include?('TradingExchangeList')
                @current_hash[@current_node] = string
              end

            when 'PrimaryExchange'
              if @stack_values.include?('TradingExchangeList')
                @current_hash[@current_node] = string
              end

            when 'XFMQ'
              if @stack_values.include?('SEDOLOfficialListingExchangeList')
                @current_hash[@current_node] = string
              end

            when 'SEDOL'
              if @stack_values.include?('SEDOLOfficialListingExchangeList')
                @current_hash[@current_node] = string
              end
            when 'Value'
              if @stack_values.include?('HistoryDetail')
                @current_hash[@current_node] = string
              end
            else
              @parsed_node[@current_node] = string
          end
        end

        def end_element name
          @stack_values.pop

          case name
            when 'GrowthOf10KNAV'
              @array = nil

            when 'SEDOLOfficialListingExchangeList'
              if @stack_values.include?('TradingExchange')
                @parsed_node['TradingExchange'][name] = @current_hash
              end

            when 'TradingExchange'
              @array = nil
              @current_hash = nil

            when 'TradingExchangeList'
              if @stack_values.include?('TradingExchange')
                @array.push @current_hash
              end

            when 'HistoryDetail'
              if @stack_values.last == 'GrowthOf10KNAV'
                @array.push @current_hash
                @current_hash = nil
              end

            when @object_node
              @data_appenders.each{|appender| appender.append(@parsed_node)}
              @parsed_node = {}
          end
        end
      end
    end
  end
end
