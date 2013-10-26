var FindService = require('./find_service');

var service = {
	findService: FindService(),

	run: function (input) {
		if (!input) throw new Error('Input invalid');
		this.input = input;
		return this.complexProcessing();
	},

	complexProcessing: function () {
		var self = this;
		var found = this.findService.run();

		return found.map(function (e) {
			return e * self.input;
		});
	}
}

module.exports = function () {
	return Object.create(service);
}