# Given two sorted arrays nums1 and nums2 with the size
# of m and n respectivley, and return the median of the two sorted
# arrays: the overall run time complexity should be O(log(min(m,n)))
from typing import List 

class Solution:

    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int] ) -> float:

        # Preprocess Step: Make the smaller of the two arrasy nums1
        m = len(nums1)
        n = len(nums2)
        if m > n:
            nums1, nums2, m, n = nums2, nums1, n, m
        
        low, high, half = 0, m-1, (m+n) // 2
        
        while True:
            # The goal is to have the lists split by two indices which cover
            # the left and right values that sit between it (e.g. [1,2 | 3, 4]) 
            # We want two find whe indices in both arrays where both indices 
            # seperate left and right values that are lA < rB and lB < rA 

            # Index will start at half the first array (0 + m - 1) // 2
            index_A = (low + high) >> 1
            # Second index should start such that sum of t
            index_B = half - index_A - 2

            print(len(nums1[0:index_A+1]))
            print(len(nums2[0:index_B+1]))
            print(index_A)
            print('\n')

            # Need to ensure sum of left halfs is (m + n + 1)//2 elements
            # ||nums1[0:(m - 1) // 2]|| + ||nums2(0, (n-1) // 2 -2)||
            # (assume m and n are even)
            # m/2 + n/2 

            # So now sum of the lower parts of the indexs is
            max_left_A  = float('-inf') if (index_A < 0) else nums1[index_A]
            min_right_A = float('inf') if (index_A + 1 >= m) else nums1[index_A + 1]
            max_left_B  = float('-inf') if (index_B < 0) else nums2[index_B]
            min_right_B = float('inf') if (index_B + 1 >= n) else nums2[index_B + 1]

            if (max_left_A <= min_right_B and max_left_B <= min_right_A):
                if ( (m+n) % 2 == 0):
                    return ((max(max_left_A, max_left_B) + min(min_right_B, min_right_A)) / 2)
                return (min(min_right_A, min_right_B))
            
            # Too far to the right
            elif(max_left_A > min_right_B): 
                high = index_A - 1

            # Too far to the left
            else: 
                low = index_A + 1


if __name__ == "__main__":
    
    a1 = [1,2,3,8,10] #11,20,30,40,41,90,91,800,1999,2000,2001]
    a2 = [4,5,6,7,20] #34,45,46,47,90,100,110,230,345,345,1000]

    solution = Solution()
    print(solution.findMedianSortedArrays(a1,a2))
