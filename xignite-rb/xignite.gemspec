# coding: utf-8
lib = File.expand_path('../lib', __FILE__)
$LOAD_PATH.unshift(lib) unless $LOAD_PATH.include?(lib)
require 'xignite/version'

Gem::Specification.new do |spec|
  spec.name          = 'xignite'
  spec.version       = Xignite::VERSION
  spec.authors       = ['Victor Paschenko', 'Valentina Grinchik', 'Sergey Privalov']
  spec.email         = ['victor@forwardlane.com', 'valentina@forwardlane.com', 'sergey@forwardlane.com']

  spec.summary       = %q{Xignite RESTful web services client implemented in Ruby}
  spec.description   = %q{Xignite RESTful web services client implemented in Ruby}
  spec.homepage      = 'https://github.com/fintechsandbox/project-sandcastle/xignite-rb'

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

  spec.add_dependency 'multi_json', '~> 3.4.0'

  spec.add_development_dependency 'bundler', '~> 1.11'
  spec.add_development_dependency 'rake', '~> 10.0'
  spec.add_development_dependency 'rspec', '~> 1.11.2'
end

