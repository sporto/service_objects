require_relative "find_service"

class ProcessService

	attr_writer :find_service

	def run(input)
		@input = input
		complex_processing
	end

	def complex_processing
		items = find_service.run
		items.map do |i|
			i * @input
		end
	end

	def find_service
		@find_service ||= FindService.new
	end

end