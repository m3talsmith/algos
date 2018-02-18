/* FindDuplicatesNaive returns any duplicates in slice of integers
 *
 * Hypothesis
 * -------
 * If we sort the integers then check the neighbor to the right, we'll know if
 * there's a duplicate. If there is a duplicate, we can append that to the
 * list of previously found duplicates, increase the index by two, and
 * continue. If there is no duplicate, we increase the index by one and
 * continue. We will have to check that the right isn't out of bounds before
 * checking the for a duplicate. If there's a duplicate, we advance the right
 * neighbor index until it no longer matches. Once we have gone through the
 * whole list we'll return the list of possible duplicates.
 */
fn find_duplicates_naive(numbers: &Vec<i32>) -> Vec<i32> {
    let mut duplicates: Vec<i32> = Vec::new();
    let mut sorted_numbers: Vec<i32> = Vec::new();
    sorted_numbers.extend_from_slice(&numbers);
    sorted_numbers.sort();

    {
        let count = numbers.len();
        let mut i = 0; // Index
        while i < count {
            let mut ni = i + 1; // Next Index
            if ni >= count {
                break;
            }
            if sorted_numbers[i] == sorted_numbers[ni] {
                duplicates.push(sorted_numbers[i]);
                while sorted_numbers[i] == sorted_numbers[ni] {
                    ni = ni + 1;
                }
            }
            i = ni;
        }
    }

    duplicates
}

fn main() {
    let n: Vec<i32> = vec![1,2,3,4,5,5,6,7,8,1];
    let k: Vec<i32> = vec![1,5];
    let d: Vec<i32> = find_duplicates_naive(&n);

    println!("Given test vector(n) {:?}", n);
    println!("Given known vector(k) {:?}", k);
    println!("When f(&Vec<i31>) -> Vec<i32> is applied to n");
    println!("Then the result will be duplicates vector(d) {:?}", d);
    println!("Which shall be equal to k.");
    println!("\nResults: d{:?} == k{:?}", d, k);

    assert_eq!(d, k);
}
