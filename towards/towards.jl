#!/usr/bin/env julia

function Towards1(n::Int64)
    a = Int64[]
    while n > 1
        push!(a, n)
        if n % 2 == 0
            n = n/2
        else
            n = n*3 + 1
        end
    end
    push!(a,n)
    return a
end
