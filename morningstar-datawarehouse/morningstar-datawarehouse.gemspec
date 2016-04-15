# coding: utf-8
lib = File.expand_path('../lib', __FILE__)
$LOAD_PATH.unshift(lib) unless $LOAD_PATH.include?(lib)
require 'morningstar/datawarehouse/version'

Gem::Specification.new do |spec|
  spec.name          = "morningstar-datawarehouse"
  spec.version       = Morningstar::Datawarehouse::VERSION
  spec.authors       = ['Victor Paschenko', 'Valentina Grinchik', 'Sergey Privalov']
  spec.email         = ['victor@forwardlane.com', 'valentina@forwardlane.com', 'sergey@forwardlane.com']

  spec.summary       = %q{Morningstar data warehouse crawler implemented in Ruby}
  spec.description   = %q{Morningstar data warehouse crawler implemented in Ruby}
  spec.homepage      = 'https://bitbucket.org/forwardlane/morningstar-datawarehouse'

  # Prevent pushing this gem to RubyGems.org by setting 'allowed_push_host', or
  # delete this section to allow pushing this gem to any host.
  if spec.respond_to?(:metadata)
    spec.metadata['allowed_push_host'] = "TODO: Set to 'http://mygemserver.com'"
  else
    raise 'RubyGems 2.0 or newer is required to protect against public gem pushes.'
  end

  spec.files         = `git ls-files -z`.split("\x0").reject { |f| f.match(%r{^(test|spec|features)/}) }
  spec.bindir        = 'bin'
  spec.executables   = spec.files.grep(%r{^bin/}) { |f| File.basename(f) }
  spec.require_paths = ['lib']

  spec.add_dependency 'aws-sdk', '~> 2'
  spec.add_dependency 'nokogiri', '~> 1.6.7.1'
  spec.add_dependency 'rubyzip', '>= 1.0.0'


  spec.add_development_dependency 'bundler', '~> 1.11'
  spec.add_development_dependency 'rake', '~> 10.0'
  spec.add_development_dependency 'rspec', '~> 3.4.0'
end
