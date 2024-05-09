def fibonacci(n)
  if n == 0
    return 0
  end

  if n == 1
    return 1
  end

  return fibonacci(n-1) + fibonacci(n-2);
end

for i in 1..50
  result = fibonacci(i)
  puts "#{i}: #{result} (#{Time.now.to_i})"
end
