ALTER TABLE `account_balances` ADD CONSTRAINT `fk_account_balances_accounts` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`account_id`) ON UPDATE CASCADE;
ALTER TABLE `account_details` ADD CONSTRAINT `fk_account_details_accounts` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`account_id`) ON UPDATE CASCADE;
ALTER TABLE `account_flags` ADD CONSTRAINT `fk_account_flags_accounts` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`account_id`) ON UPDATE CASCADE;
ALTER TABLE `accounts` ADD CONSTRAINT `fk_accounts_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON UPDATE CASCADE;
ALTER TABLE `banners` ADD CONSTRAINT `fk_banners_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON UPDATE CASCADE;
ALTER TABLE `debit_card_design` ADD CONSTRAINT `fk_debit_card_design_debit_cards` FOREIGN KEY (`card_id`) REFERENCES `debit_cards` (`card_id`) ON UPDATE CASCADE;
ALTER TABLE `debit_card_details` ADD CONSTRAINT `fk_debit_card_details_debit_cards` FOREIGN KEY (`card_id`) REFERENCES `debit_cards` (`card_id`) ON UPDATE CASCADE;
ALTER TABLE `debit_card_status` ADD CONSTRAINT `fk_debit_card_status_debit_cards` FOREIGN KEY (`card_id`) REFERENCES `debit_cards` (`card_id`) ON UPDATE CASCADE;
ALTER TABLE `debit_cards` ADD CONSTRAINT `fk_debit_cards_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON UPDATE CASCADE;
ALTER TABLE `transactions` ADD CONSTRAINT `fk_transactions_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON UPDATE CASCADE;
ALTER TABLE `user_greetings` ADD CONSTRAINT `fk_user_greetings_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON UPDATE CASCADE;