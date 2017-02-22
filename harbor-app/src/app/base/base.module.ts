import { NgModule } from '@angular/core';
import { SharedModule } from '../shared/shared.module';

import { DashboardModule } from '../dashboard/dashboard.module';
import { ProjectModule } from '../project/project.module';
import { UserModule } from '../user/user.module';
import { AccountModule } from '../account/account.module';

import { NavigatorComponent } from './navigator/navigator.component';
import { GlobalSearchComponent } from './global-search/global-search.component';
import { FooterComponent } from './footer/footer.component';
import { HarborShellComponent } from './harbor-shell/harbor-shell.component';
import { SearchResultComponent } from './global-search/search-result.component';

import { BaseRoutingModule } from './base-routing.module';

@NgModule({
  imports: [
    SharedModule,
    DashboardModule,
    ProjectModule,
    UserModule,
    BaseRoutingModule,
    AccountModule
  ],
  declarations: [
    NavigatorComponent,
    GlobalSearchComponent,
    FooterComponent,
    HarborShellComponent,
    SearchResultComponent
  ],
  exports: [ HarborShellComponent ]
})
export class BaseModule {

}