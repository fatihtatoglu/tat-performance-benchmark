local function fibonacci(n)
    if n == 0 then
        return 0
    end

    if n == 1 then
        return 1
    end

    return fibonacci(n - 1) + fibonacci(n - 2)
end

for i = 1, 98, 1 do
    local result = fibonacci(i)
    print(i, ": ", result)
end
