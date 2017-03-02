module Xignite
  class MasterDataService < Xignite::BaseService

    def base_url
      'http://globalmaster.xignite.com/xglobalmaster'
    end

    def get_master_by_identifier(params, &block)
      default_params = {
          IdentifierType: IDENTIFIER_TYPE,
          Identifier: '',
          StartDate: '',    # in format "%m/%d/%Y"
          EndDate: ''       # in format "%m/%d/%Y"
      }
      send_request('GetMasterByIdentifier', default_params.merge(params)){|data| block.call(data)}
    end

    def get_master_by_identifiers(params, &block)
      default_params = {
          IdentifierType: IDENTIFIER_TYPE,
          Identifiers: '',
          AsOfDate: '',    # in format "%m/%d/%Y"
      }
      send_request('GetMasterByIdentifiers', default_params.merge(params)){|data| block.call(data)}
    end

    def list_exchanges(params, &block)
      default_params={}
      send_request('ListExchanges', default_params.merge(params)){|data| block.call(data)}
    end

    def list_mic_to_legacy_exchange(params, &block)
      default_params={}
      send_request('ListMICToLegacyExchange', default_params.merge(params)){|data| block.call(data)}
    end

    def get_master_by_exchange_changes(params, &block)
      default_params={
          Exchange: '',
          StartSymbol:'',
          EndSymbol:'',
          InstrumentClass:'All',
          StartModifiedDate:'',
          EndModifiedDate:''
      }
      send_request('GetMasterByExchangeChanges', default_params.merge(params)){|data| block.call(data)}
    end

    def get_master_by_exchange(params, &block)

      default_params={
          Exchange: '',
          StartSymbol:'',
          EndSymbol:'',
          InstrumentClass:'All',
          AsOfDate: ''
      }
      send_request('GetMasterByExchange', default_params.merge(params)){|data| block.call(data)}
    end

    def get_instruments(params, &block)
      default_params={
          IncludeRelated: 'Securities',
          IdentifierType: IDENTIFIER_TYPE,
          Identifiers: '',
          AsOfDate: ''
      }
      send_request('GetInstruments', default_params.merge(params)){|data| block.call(data)}
    end

    def get_instrument(params, &block)
      default_params={
          IncludeRelated: 'None',
          IdentifierType: IDENTIFIER_TYPE,
          Identifier: '',
          StartDate: '',
          EndDate: ''
      }
      send_request('GetInstrument', default_params.merge(params)){|data| block.call(data)}
    end

  end
end
