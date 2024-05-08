(function () {
    for (let i = 0; i <= 98; i++) {
        var result = fibonacci(i)
        console.log(i, ":", result);
    }

    function fibonacci(n) {
        if (n === 0) return 0;
        if (n === 1) return 1;

        return fibonacci(n - 1) + fibonacci(n - 2);
    }
})();