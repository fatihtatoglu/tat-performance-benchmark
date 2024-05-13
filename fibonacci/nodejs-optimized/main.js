(function () {
    for (let i = 0; i <= 50; i++) {
        var result = fibonacci(i)

        var epoch = Date.now();

        console.log(i, ":", result, "(", epoch, ")");
    }

    function fibonacci(n) {
        if (n === 0) return 0;
        if (n === 1) return 1;

        return fibonacci(n - 1) + fibonacci(n - 2);
    }
})();