#include <iostream>
#include <list>
#include <vector>
#include <assert.h>

std::list<int> findDuplicatesNaive(std::list<int> numbers) {
    int index = 0;
    int nextIndex;
    int count = numbers.size();
    std::list<int> duplicates;

    numbers.sort();

    std::vector<int> sortedNumbers {
        std::begin(numbers),
        std::end(numbers)
    };

    int currentValue;
    int nextValue;

    while (index < count) {
        nextIndex = index + 1;
        if (nextIndex >= count) {
            index = nextIndex;
            break;
        }

        currentValue = sortedNumbers[index];
        nextValue = sortedNumbers[nextIndex];

        std::cout
            << "currentValue: " << currentValue
            << ", nextValue: "  << nextValue
            << std::endl;

        if (currentValue == nextValue) {
            std::cout << "duplicate found: " << currentValue << std::endl;
            duplicates.push_back(currentValue);
            while (sortedNumbers[index] == sortedNumbers[nextIndex]) {
                std::cout
                    << "sortedNumbers[" << index << "]: " << sortedNumbers[index]
                    << ", sortedNumbers[" << nextIndex << "]: " << sortedNumbers[nextIndex]
                    << std::endl;
                nextIndex += 1;
            }
        }
        index = nextIndex;
    }

    return duplicates;
}

int main() {
    std::list<int> n1 = {1,2,3,4,5,5,6,7,8,1};
    std::list<int> t1 = findDuplicatesNaive(n1);
    std::list<int> e1 = {1,5};
    assert(t1 == e1);
    return 0;
}
