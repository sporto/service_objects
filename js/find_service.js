var service = {
	run: function () {
		return [1, 2];
	}
}

module.exports = function () {
	return Object.create(service);
}