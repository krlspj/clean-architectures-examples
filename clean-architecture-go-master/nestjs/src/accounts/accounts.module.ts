import { Module } from '@nestjs/common';
import { AccountsService } from './accounts.service';
import { AccountsController } from './accounts.controller';
import {SequelizeModule} from "@nestjs/sequelize";
import {Account} from "./entities/account.entity";
import { AccountStorageService } from './account-storage/account-storage.service';
import {TokenGuard} from "./token.guard";

@Module({
  imports: [SequelizeModule.forFeature([Account])],
  controllers: [AccountsController],
  providers: [AccountsService, AccountStorageService, TokenGuard],
  exports: [AccountStorageService]
})
export class AccountsModule {}
