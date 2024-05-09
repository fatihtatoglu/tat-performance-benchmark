for $i (0 .. 50) {
    $result = fibonacci($i);

    my $epoch = time;	

    print $i, ": ", $result, "(",$epoch,")\r\n";
}

sub fibonacci{
    my ($n) = @_;
    if ($n == 0) {
        return 0;
    }
    
    if ($n == 1) {
        return 1;
    }

    return fibonacci($n-1) + fibonacci($n-2);
}