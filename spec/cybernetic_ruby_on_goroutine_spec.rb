# frozen_string_literal: true

RSpec.describe CyberneticRubyOnGoroutine do
  it "has a version number" do
    expect(CyberneticRubyOnGoroutine::VERSION).not_to be nil
  end

  it "can sum numbers" do
    expect(CyberneticRubyOnGoroutine.sum(1, 2)).to eq(3)
  end
end
