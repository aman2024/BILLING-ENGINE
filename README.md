# BILLING-ENGINE-SYSTEM

Assumptions:
1. Authorization contains userId/adminId
2. Authorization is of type NoAuth with only value userId/adminId. Skipping encoding/decoding part
3. View api is returning all result of users (Not doing in pagination manner)
4. Admin blindly approves the loan without seeing his past history/balance
5. 


Steps to run the service
1. Install are the pre-requisite written in requirement.txt file
2. Set MYSQL username, host, port, password in .env.test file
3. Create table in MYSQL. DDL cmd in shared below
4. Run either of the cmds on terminal to run the service 
    a. ENV=dev go run .
    b. sh start.sh



MYSQL table Schema

CREATE TABLE `billing_info` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` varchar(45) NOT NULL,
  `amount` int NOT NULL,
  `initial_term` int NOT NULL,
  `rate` int NOT NULL,
  `repayment_amount` int NOT NULL,
  `status` varchar(45) NOT NULL,
  `admin_approver_id` varchar(45) DEFAULT NULL,
  `amount_paid` float DEFAULT '0',
  `is_delinquent` bool DEFAULT false,
  `total_terms` int NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `approved_at` datetime,
  PRIMARY KEY (`id`)
)

CREATE TABLE `repayment_info` (
  `id` int NOT NULL AUTO_INCREMENT,
  `loan_id` int NOT NULL,
  `repayment_amount` float NOT NULL,
  `term_no` int NOT NULL,
  `repayment_date` datetime NOT NULL,
  `status` varchar(45) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
)


Postman Collection: Link is shared in mail. Can't share here due to access violation since it has access key

There are 4 api's: 
All require Authorization(in Headers)(Key: Authorization, Value:user_id/admin_id)
1. Create Loan: This api is used to create loan
2. Approve Loan: This api is used to approve loan(by admin)
3. View Loan: This api is used to view loan including OutstandingBalance and IsDelinquent
4. Add Repayment: This api is used to add repayment

There are 2 cron's:
Both the cron's are running at midnight
1. UpdateRepaymentStatusFromLockedToPending: 
  a. This is used to update the status of the repayment to PENDING state for the term which needs to enabled for this week
2. UpdateRepaymentStatusFromPendingToSkipped: 
  a. This is used to update the status of the repayment to SKIPPED state for the term which is not paid by the user. 
  b. Also it creates a new entry of repayment for the corresponding week in last. 
  c. Also checks whether this loan is Delinquent or not. If delinquent, it will update it in billing_info table

  Test Cases:
  Very basic unit test cases are written to ensure individual parts of the software function correctly and reliably.