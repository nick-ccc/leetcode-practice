class Solution(object): 
    def findKthLargest(self,nums,k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: int
        
        Thought process: 
        Quick Sort - a divide and conquer algorithim that picks
        an element as a pivor and paritions the given array around 
        the selected picked pivot by placing the pivot in its correct 
        position in the sorted array 

        You can use quick sort until the kth largest element is 
        selected as a pivot 
        """
        
        def swap(arr, i, j):
            arr[i], arr[j] = arr[j], arr[i]
       
        def partition(arr, low, high):
            # Start with a 'random' pivot, here we just use high
            pivot = arr[high]
            i     = low

            for j in range(low,high):
                if arr[j] <= pivot:
                    swap(arr, i, j)
                    i += 1
            swap(arr, i, high)
            return i  
        
        def quick_select(arr, low, high, k):
            if low <= high: 
                #print(arr)
                pi = partition(arr,low,high)

                if(pi == k):
                    return arr[pi]
                elif pi < k:
                    return quick_select(arr, pi + 1, high, k)
                else:
                    return quick_select(arr,low, pi - 1, k) 

        # To get the kth largest convert k into len(arr) - k
        # k = 1 -> maps to len(arr) - 1 i.e. the largest value
        k = len(nums) - k
        return quick_select(nums,0, len(nums) - 1, k)

if __name__ == "__main__":
    sol = Solution()
    arr = [0,2,5,6,7,9,5,1,4,8]
    print(len(arr))
    for i in range(1, len(arr) + 1):
        print(sol.findKthLargest(arr,i))
        arr = [0,2,5,6,7,9,5,1,4,8]
