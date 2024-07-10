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





