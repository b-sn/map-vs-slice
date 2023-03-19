#!/usr/bin/env perl

use Modern::Perl;

# Function returns the number of each character in the string
sub count_chars_hash {
    my $chars = shift;
    my %count;
    $count{ord($_)}++ for @$chars;
    return \%count;
}

sub count_chars_array {
    my $chars = shift;
    my @count;
    $count[ord($_)]++ for @$chars;
    return \@count;
}

# Read the file ./1984.txt to string variable
my $file = do { local $/; open my $fh, '<', './1984.txt'; <$fh> };

my @chars = split //, $file;

# Benchmark the functions count_chars_hash and count_chars_array
use Benchmark qw(cmpthese);
# my $cntHash, $cntArray;
cmpthese(1000, {
    'count_chars_hash' => sub { count_chars_hash(\@chars) },
    'count_chars_array' => sub { count_chars_array(\@chars) },
});
# print $cntArray;
# print $cntHash;