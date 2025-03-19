# frozen_string_literal: true

require_relative "lib/cybernetic_ruby_on_goroutine/version"

Gem::Specification.new do |spec|
  spec.name = "cybernetic_ruby_on_goroutine"
  spec.version = CyberneticRubyOnGoroutine::VERSION
  spec.authors = ["Isamu Mogi"]
  spec.email = ["isamu@leafytree.jp"]

  spec.summary = "サイバネティックなRuby on Goroutine"
  spec.description = "正気でいられるなんで運がイイぜ"
  spec.homepage = "https://github.com/saturday06/cybernetic_ruby_on_goroutine"
  spec.license = "MIT"
  spec.required_ruby_version = ">= 3.1.0"

  spec.metadata["allowed_push_host"] = "https://example.com"

  spec.metadata["homepage_uri"] = spec.homepage
  spec.metadata["source_code_uri"] = "https://github.com/saturday06/cybernetic_ruby_on_goroutine"
  spec.metadata["changelog_uri"] = "https://github.com/saturday06/cybernetic_ruby_on_goroutine/blob/main/CHANGELOG.md"

  gemspec = File.basename(__FILE__)
  spec.files = IO.popen(%w[git ls-files -z], chdir: __dir__, err: IO::NULL) do |ls|
    ls.readlines("\x0", chomp: true).reject do |f|
      (f == gemspec) ||
        f.start_with?(*%w[bin/ test/ spec/ features/ .git .github appveyor Gemfile])
    end
  end
  spec.bindir = "exe"
  spec.executables = spec.files.grep(%r{\Aexe/}) { |f| File.basename(f) }
  spec.require_paths = ["lib"]
end
