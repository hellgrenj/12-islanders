# 12 islanders 

Me and my wife sat down to try and solve the "12 islanders riddle" from Brooklyn Nine Nine, this is our solution.  

## the riddle

*“There are 12 men on an island. 11 of them weigh the same and one of them is either slightly heavier or slightly lighter. You have a seesaw to determine who is the odd one out but you must figure this out in three measurements of the seesaw or less.”*  

## our solution as code (simulation + test)
(You need Go installed)  
run a simulation: ``go run .``  
test all the cases: ``go test . -v``  

## our solution described

1 split the 12 islanders into three groups  
seesaw: L1 L2 L3 L4     R1 R2 R3 R4     
sideline: S1 S2 S3 S4  
(named after their starting position, L1 = Left 1, R1 = Right 1, S1 = Side 1 etc..)  

(seesaw measurement number 1)

2.A if balanced

We now know that the odd islander are amongst **S1 S2 S3 S4**  
seesaw: L1 L2 L3 L4     R1 R2 R3 R4     (balanced)  
sideline: **S1 S2 S3 S4**    

we now compare S1 S2 S3 with R1 R2 R3.

(seesaw measurement number 2)

2.A.1 if balanced   
seesaw: S1 S2 S3    R1 R2 R3 (balanced)  
sideline: **S4**  L1 L2 L3 L4 R4  
we now know that the odd islander is **S4**. At this point we just weight S4 against a 
neutral islander to figure out if it is heavy or light.  
(seesaw measurement number 3)

2.A.2 if left is heavy   
seesaw: **S1 S2 S3** (heavy)      R1 R2 R3  
sideline: L1 L2 L3 S4 L4 R4    
we now know that the heavier islander are amongst **S1 S2 S3**.  
we weigh S1 against S2; if balanced it is S3 - else it is the heavy of S1 and S2.  
(seesaw measurement number 3)

2.A.3 if right is heavy  
seesaw: **S1 S2 S3**  R1 R2 R3 (heavy)   
sideline: L1 L2 L3 S4 L4 R4  

we now know that the lighter islander are amongst **S1 S2 S3**.  
we weigh S1 against S2. if balanced it is S3 - else it is the lighter of S1 and S2.  
(seesaw measurement number 3)

2.B if unbalanced  

*(first make sure you have the heavy side on the left...)*  

seesaw: L1 L2 L3 L4 (heavy)      R1 R2 R3 R4  
sideline: S1 S2 S3 S4  

now switch R1 R2 R3 and S2 S3 S4 AND switch R1 AND L1. 

(seesaw measurement number 2)

2.B.1 if left heavy  
seesaw: R1 **L2 L3 L4** (heavy)   L1 S2 S3 S4  
sideline: S1 R2 R3 R4  

we know there is a heavy islander amongst **L2 L3 L4**  
now compare L2 and L3. if balanced it is L4 - else it is the heavier of L2 and L3.  
(seesaw measurement number 3)


2.B.2 if balanced  
seesaw: R1 L2 L3 L4    L1 S2 S3 S4  (balanced)  
sideline: S1 **R2 R3 R4** 

we know that the light islander is amongst **R2 R3 R4**  
now compare R2 and R3. if balanced it is R4 - else it is the lighter of R3 and R2   
(seesaw measurement number 3)


2.B.3 if right heavy   
seesaw: **R1** L2 L3 L4    **L1** S2 S3 S4 (heavy)  
sideline: S1 R2 R3 R4  

we know it is either **R1** or **L1**.  
now compare R1 against a neutral islander (S1 for example).  
If balanced it is the heavy L1 else if unbalanced it is light R1.   
(seesaw measurement number 3).   

