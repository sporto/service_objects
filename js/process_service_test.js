// var assert = require("assert");
var Service = require('./process_service');
var expect = require('chai').expect;

describe('ProcessService', function(){

	var serv;

	beforeEach(function () {
		serv = Service();
	});

	describe('.run', function(){
		it('uses the defaults', function(){
			var res = serv.run(2);
			expect(res).to.eql([2, 4]);
		});

		it('can take a fake', function () {
			var fakeFindService = {
				run: function () {
					return [1, 2, 3];
				}
			}
			serv.findService = fakeFindService;
			var res = serv.run(2);
			expect(res).to.eql([2, 4, 6]);
		});
	});
})