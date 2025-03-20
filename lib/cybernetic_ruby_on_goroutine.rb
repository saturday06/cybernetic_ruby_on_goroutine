# frozen_string_literal: true

require_relative "cybernetic_ruby_on_goroutine/version"

module CyberneticRubyOnGoroutine
  class Error < StandardError; end

  # これは世紀末
  class EndOfTheCentury
    def process
      puts "ok"
    end
  end
end
