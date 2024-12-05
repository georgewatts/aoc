d = {(i,j):c for (i,l) in enumerate(open("input2.txt")) for (j,c) in enumerate(l.strip())}
delta = ((1,2,3,0,0,0), (-1,-2,-3,0,0,0), (0,0,0,1,2,3), (0,0,0,-1,-2,-3), (1,2,3,1,2,3), (-1,-2,-3,1,2,3), (1,2,3,-1,-2,-3), (-1,-2,-3,-1,-2,-3))
print(sum(d.get((i+u,j+x))=="M" and d.get((i+v,j+y))=="A" and d.get((i+w,j+z))=="S" for (u,v,w,x,y,z) in delta for (i,j) in d if d[(i,j)] == "X"))
print(sum(1 for (i, j) in d if d[(i,j)] == "A" and {d.get((i-1,j-1)), d.get((i+1,j+1))} == {d.get((i-1,j+1)), d.get((i+1,j-1))} == {"M","S"}))
