#include <iostream>
#include <string>
#include "704-binary-search.h"
#define assert(x, msg) \
    if(!(x)) { fprintf(stderr, "Fatal error: %s\n", msg); abort(); }

int main()
{
    Solution s;
    // {
    //     vector<int> testcase1 = {-1,0,3,5,9,12};
    //     int target = 9;
    //     assert((void("asasd"), s.search(testcase1, target)==4));
    // }
    {
        vector<int> testcase2 = {-1,0,3,5,9,12};
        int target = 9;
        int expected = 4;
        assert(s.search(testcase2, target)==expected, "Fail");
    }


    return 0;
}