from Vector import Vector

if __name__=="__main__":
    vec =Vector([5,2])
    print(vec)
    print(len(vec))
    print(vec[0])
    vec1= Vector([1,2])
    print(vec1+vec)
    print(2*(vec1+vec))
    print((2*vec)+(2*vec1))

    print(vec.normalize())
    print(vec.normalize().norm())

    zero2=Vector.zero(2)
    
    try:        
        print(zero2.normalize().norm())
    except ZeroDivisionError:
        print(f"Cannot normalize zero vector {zero2}")
