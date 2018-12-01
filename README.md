# Challenge 3 - Match a person with Levenshtein Distance

## Deadline: yyyy-MM-dd HH:mm:ss SSS

## Description

You are tasked by the business owners on the Marketing team to improve the matching process of a Lead in the system.  Currently, a person (lead) is found in the system and considered a "match" if the input (search criteria) `First and Last Name` is equal to the system's `First and Last Name`.

We need to relax our matching logic slightly and implement fuzzy (approximate) string matching on the first name and last name.

We will be using the `Levenshtein Distance` to implement this.

## Requirements

* Read the file, `leads.csv` which is located at the project root and use as your data store.
* The application should accept a `first name` and `last name` (or a data structure representing each) as INPUT.
* The INPUT will be compared against the _Leads_ in the data store (`leads.csv`) using _Levenshtein Distance_.
* The best match's `row ID` will be returned with the `First Name` and the `Last Name` of the respective match (again, feel free to return 3 values if your language supports it; or a data structure representing the 3 values).

### Matching

The Marketing team has come up with the following match requirements using this fuzzy matching algorithm.

**It is possible that there could be multiple matches that meet the below criteria.  You should return the _best_ match (lowest Levenshtein Distance).  If there are 2+ matches with the _same_ distance, then simply return one of them.**

* The Levenshtein Distance for `First` Name should be *no larger* than *3*.
* The Levenshtein Distance for `Last` Name should be *no larger* than *1*.
* The total Levenshtein Distance between `First` _AND_ `Last` name should be *no larger* than *3*.
* The required Levenshtein Distance for `Last Name`should be configurable so that the business can increase/decrease the tolerance as needed.
**NOTE** if the `Last Name` tolerance is configurable, you will need to update the `Total`'s tolerance appropriately.
* Case sensitivity should be ignored and should not contribute to the Levenshtein Distance (ie. `e` == `E`).

## Assumptions

* `leads.csv` is normalized to UPPERCASE
* There WILL NOT be any surprises in the data, such as First/Last Name in the `FirstName` column and a `null` `LastName` column.
* The only possible characters within `leads.csv` are `A-Z` or `'`, plus `0-9` for the row ID.

## Additional Instructions

* Use a language of your choice.
* NO 3rd party libraries or frameworks.
* Include your solution in a subdirectory that is named appropriately to determine who submitted it.
* Feel free to add names to `leads.csv`.  The initial data in the file is meant to provide some examples of input data.  
* After submissions are in, a more extensive input file will be created and used for the testing/benchmarking.
