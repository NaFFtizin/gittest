i = 1
puts "Посмотри в папку с программой"
while i == 1
	if !File.file?("XAXAXA")
		f = File.new("XAXAXA", "w+")
		f.close
	end
end