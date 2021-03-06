
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function ColorDialog_Create(AOwner: TComponent): TColorDialog; extdecl;
begin
  Result :=  TColorDialog.Create(AOwner);
end;

procedure ColorDialog_Free(AObj: TColorDialog); extdecl;
begin
  AObj.Free;
end;

function ColorDialog_Execute(AObj: TColorDialog): LongBool; extdecl;
begin
  Result :=  AObj.Execute;
end;

function ColorDialog_FindComponent(AObj: TColorDialog; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function ColorDialog_GetNamePath(AObj: TColorDialog): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function ColorDialog_HasParent(AObj: TColorDialog): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure ColorDialog_Assign(AObj: TColorDialog; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function ColorDialog_ClassType(AObj: TColorDialog): TClass; extdecl;
begin
  Result :=  AObj.ClassType;
end;

function ColorDialog_ClassName(AObj: TColorDialog): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function ColorDialog_InstanceSize(AObj: TColorDialog): Integer; extdecl;
begin
  Result :=  AObj.InstanceSize;
end;

function ColorDialog_InheritsFrom(AObj: TColorDialog; AClass: TClass): LongBool; extdecl;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function ColorDialog_Equals(AObj: TColorDialog; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function ColorDialog_GetHashCode(AObj: TColorDialog): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function ColorDialog_ToString(AObj: TColorDialog): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function ColorDialog_GetColor(AObj: TColorDialog): TColor; extdecl;
begin
  Result :=  AObj.Color;
end;

procedure ColorDialog_SetColor(AObj: TColorDialog; AValue: TColor); extdecl;
begin
  AObj.Color := AValue;
end;

function ColorDialog_GetHandle(AObj: TColorDialog): HWND; extdecl;
begin
  Result :=  AObj.Handle;
end;

procedure ColorDialog_SetOnClose(AObj: TColorDialog; AEventId: NativeUInt); extdecl;
begin
  {$DEFINE EventName := OnClose}
  {$DEFINE EventMethodName := OnClose}
  EventMethodCode
end;

procedure ColorDialog_SetOnShow(AObj: TColorDialog; AEventId: NativeUInt); extdecl;
begin
  {$DEFINE EventName := OnShow}
  {$DEFINE EventMethodName := OnShow}
  EventMethodCode
end;

function ColorDialog_GetComponentCount(AObj: TColorDialog): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function ColorDialog_GetComponentIndex(AObj: TColorDialog): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure ColorDialog_SetComponentIndex(AObj: TColorDialog; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function ColorDialog_GetOwner(AObj: TColorDialog): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function ColorDialog_GetName(AObj: TColorDialog): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure ColorDialog_SetName(AObj: TColorDialog; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function ColorDialog_GetTag(AObj: TColorDialog): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure ColorDialog_SetTag(AObj: TColorDialog; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function ColorDialog_GetComponents(AObj: TColorDialog; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;

function ColorDialog_StaticClassType: TClass; extdecl;
begin
  Result := TColorDialog;
end;

exports
  ColorDialog_Create,
  ColorDialog_Free,
  ColorDialog_Execute,
  ColorDialog_FindComponent,
  ColorDialog_GetNamePath,
  ColorDialog_HasParent,
  ColorDialog_Assign,
  ColorDialog_ClassType,
  ColorDialog_ClassName,
  ColorDialog_InstanceSize,
  ColorDialog_InheritsFrom,
  ColorDialog_Equals,
  ColorDialog_GetHashCode,
  ColorDialog_ToString,
  ColorDialog_GetColor,
  ColorDialog_SetColor,
  ColorDialog_GetHandle,
  ColorDialog_SetOnClose,
  ColorDialog_SetOnShow,
  ColorDialog_GetComponentCount,
  ColorDialog_GetComponentIndex,
  ColorDialog_SetComponentIndex,
  ColorDialog_GetOwner,
  ColorDialog_GetName,
  ColorDialog_SetName,
  ColorDialog_GetTag,
  ColorDialog_SetTag,
  ColorDialog_GetComponents,
  ColorDialog_StaticClassType;

