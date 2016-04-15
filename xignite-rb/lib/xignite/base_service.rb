require 'net/http'
require 'json'
require 'yaml'
require 'timeout'

module Xignite

  class RequestError < ::RuntimeError

  end

  class BaseService

    IDENTIFIER_TYPE = 'ISIN'

    BASE_REQUEST_TYPE = {
        json: 'json',
        xml: 'xml',
        csv: 'csv'
    }

    def base_url
      'https://xignite.com/'
    end

    def token
      Xignite::Configuration.options.apitoken
    end

    def request_timeout
      Xignite::Configuration.options.request_timeout || 10
    end

    protected

    # Send request
    # parameters:
    #   action(required) - api action name 'GetInstrument'
    #   options additions params
    #   callbackBlock calback after request was done
    # after getting response there will be possibility to process data like a block
    # for example BaseService.send_request(action, request_type, options={}, &callbackBlock){|data| doStuffWithData(data)}
    def send_request(action, options, &callbackBlock)
      url = construct_request_uri(action, :json, options)
      begin
        response = nil
        Timeout::timeout(request_timeout) {
          response = Net::HTTP.get_response(URI.parse(url))
        }
        callbackBlock.call(JSON.parse(response.body))
      rescue Exception => e
        raise Xignite::RequestError, "Connection error #{url}\n#{e.message}"
      end
    end

    private

    # process options to create url where send data
    # parameters:
    #   action(required) - api action name 'GetInstrument'
    #   type (required) - request type (:json, :xml, :csv)
    #   options additions params
    def construct_request_uri(action, type, options={})
      base_url + ".#{BASE_REQUEST_TYPE[type]}/#{action}" + construct_params_from_options(options) + "&_token=" + token
    end

    # construct url by parameters
    def construct_params_from_options(options={})
      url = ""
      options.keys.each_with_index do |key, index|
        index == 0 ? url = "?" : url += '&'
        url+="#{key}=#{options[key]}"
      end
      url
    end

  end
end
