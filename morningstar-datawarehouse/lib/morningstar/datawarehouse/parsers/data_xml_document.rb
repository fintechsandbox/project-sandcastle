require 'nokogiri'

module Morningstar
  module Datawarehouse
    module Parsers
      class DataXmlDocument < Nokogiri::XML::SAX::Document

        attr_writer :data_appenders

        def initialize
          @object_node = 'InvestmentVehicle'

          @stack_values = []
          @parsed_node = {}

          @data_appenders = []
          @deleted_shares = []
          @under_review_shares = []
          @updated_shares = []
          @universe = nil

          @current_node = nil
          @current_value = nil
          @current_hash = nil
        end

        def with_data_appenders(appenders=[])
          @data_appenders += appenders
          self
        end

        def with_universe(universe)
          @universe = universe
          self
        end

        def with_deleted_shares(shares=[])
          (@deleted_shares += shares) if shares
          self
        end

        def with_under_review_shares(shares=[])
          (@under_review_shares += shares) if shares
          self
        end

        def with_updated_shares(shares=[])
          (@updated_shares += shares) if shares
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


        def close

          unless @deleted_shares.empty?
            @data_appenders.each{|appender| appender.deleted_shares(@deleted_shares, @universe)}
          end

          unless @under_review_shares.empty?
            @data_appenders.each{|appender| appender.under_review_shares(@under_review_shares, @universe)}
          end

          unless @updated_shares.empty?
            @data_appenders.each{|appender| appender.updated_shares(@updated_shares, @universe)}
          end
        end

      end
    end
  end
end
