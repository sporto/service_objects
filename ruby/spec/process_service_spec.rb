require 'spec_helper'
require_relative '../process_service'

describe ProcessService do

  let(:serv) { ProcessService.new }
  let(:fake_find_service) { double.as_null_object }

  it "uses the defaults" do
    res = serv.run(2)
    expect(res).to eq([2,4])
  end

  it "can take a fake" do
    fake_find_service.stub(:run).and_return([1,2,3])
    serv.find_service = fake_find_service
    res = serv.run(2)
    expect(res).to eq([2,4,6])
  end

end