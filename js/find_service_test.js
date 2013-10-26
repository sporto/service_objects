// var assert = require("assert");
var Service = require('./find_service');
var expect = require('chai').expect;

describe('FindService', function(){

	var serv;

	beforeEach(function () {
		serv = Service();
	});

	describe('.run', function(){
		it('returns an array', function(){
			var res = serv.run();
			expect(res).to.eql([1, 2]);
		})
	});
})